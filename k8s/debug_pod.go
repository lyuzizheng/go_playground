package k8s

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"

	"net/http"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	informerAppsV1 "k8s.io/client-go/informers/apps/v1"
	informerCoreV1 "k8s.io/client-go/informers/core/v1"

	"github.com/gin-contrib/cors"
	"k8s.io/client-go/informers"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	Clientset          kubernetes.Interface
	DeploymentInformer informerAppsV1.DeploymentInformer
	PodInformer        informerCoreV1.PodInformer
	wg                 sync.WaitGroup
)

type deploymentParams struct {
	Namespace string `json:"namespace" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

func init() {
	wg.Add(2)
	_, err := genDeploymentInformer(&wg)
	if err != nil {
		panic(err)
	}
	_, err = genPodInformer(&wg)
	if err != nil {
		panic(err)
	}
	wg.Wait()
	_, err = genClientSet()
	if err != nil {
		panic(err)
	}
	go deleteExpiredDebugPodLoop()
}

func main() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowOrigins:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))
	r.POST("/debugPod", createDebugPodController)
	r.GET("/deployment", getDeploymentController)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func genDeploymentInformer(wg *sync.WaitGroup) (deploymentInformer informerAppsV1.DeploymentInformer, err error) {
	kubeconfig, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*10)
	deploymentInformer = informerFactory.Apps().V1().Deployments()
	_, err = deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) {},
		UpdateFunc: func(old, new interface{}) {},
		DeleteFunc: func(obj interface{}) {},
	})
	if err != nil {
		return
	}
	DeploymentInformer = deploymentInformer
	go func() {
		stop := make(chan struct{})
		defer close(stop)
		informerFactory.Start(stop)
		informerFactory.WaitForCacheSync(stop)
		wg.Done()
		select {}
	}()
	return
}

func genPodInformer(wg *sync.WaitGroup) (podInformer informerCoreV1.PodInformer, err error) {
	kubeconfig, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*10)
	podInformer = informerFactory.Core().V1().Pods()
	_, err = podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) {},
		UpdateFunc: func(old, new interface{}) {},
		DeleteFunc: func(obj interface{}) {},
	})
	if err != nil {
		return
	}
	PodInformer = podInformer
	go func() {
		stop := make(chan struct{})
		defer close(stop)
		informerFactory.Start(stop)
		informerFactory.WaitForCacheSync(stop)
		wg.Done()
		select {}
	}()
	return
}

func genClientSet() (clientset kubernetes.Interface, err error) {
	kubeconfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	clientset, err = kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	Clientset = clientset
	return
}

func createDebugPodController(c *gin.Context) {
	var params deploymentParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不能为空"})
		return
	}
	err := createDebugPod(params, Clientset, DeploymentInformer, PodInformer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func createDebugPod(params deploymentParams, clientSet kubernetes.Interface, deploymentInformer informerAppsV1.DeploymentInformer, podInformer informerCoreV1.PodInformer) (err error) {
	deployment, err := getDeployment(params, deploymentInformer)
	if err != nil {
		return err
	}
	exist, err := getDebugPod(params, podInformer)
	if err != nil && !errors.IsNotFound(err) {
		fmt.Println("createDebugPod getDebugPod error", err)
		return err
	}
	if exist {
		err = deleteDebugPod(params, clientSet)
		if err != nil {
			return err
		}
		// wait for delete
		for {
			_, err = podInformer.Lister().Pods(params.Namespace).Get(fmt.Sprintf("%v-debug", params.Name))
			if err != nil && errors.IsNotFound(err) {
				break
			}
			time.Sleep(time.Second * 2)
		}
	}
	pod := genPod(*deployment)
	_, err = clientSet.CoreV1().Pods(params.Namespace).Create(context.Background(), &pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func getDeployment(params deploymentParams, deploymentInformer informerAppsV1.DeploymentInformer) (deployment *appsv1.Deployment, err error) {
	deployment, err = deploymentInformer.Lister().Deployments(params.Namespace).Get(params.Name)
	if err != nil {
		return
	}
	return
}

func getDebugPod(params deploymentParams, podInformer informerCoreV1.PodInformer) (exist bool, err error) {
	debugPodName := fmt.Sprintf("%v-debug", params.Name)
	_, err = podInformer.Lister().Pods(params.Namespace).Get(debugPodName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func deleteDebugPod(params deploymentParams, clientSet kubernetes.Interface) (err error) {
	debugPodName := fmt.Sprintf("%v-debug", params.Name)
	err = clientSet.CoreV1().Pods(params.Namespace).Delete(context.TODO(), debugPodName, metav1.DeleteOptions{})
	if err != nil {
		return
	}
	return nil
}

func genPod(deployment appsv1.Deployment) (pod v1.Pod) {
	pod = v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%v-debug", deployment.Name),
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":     deployment.Name,
				"debug":   "true",
				"version": "debug",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            fmt.Sprintf("%v-debug", deployment.Name),
					Image:           deployment.Spec.Template.Spec.Containers[0].Image,
					ImagePullPolicy: v1.PullIfNotPresent,
					Command:         deployment.Spec.Template.Spec.Containers[0].Command,
					Env:             deployment.Spec.Template.Spec.Containers[0].Env,
					Resources:       deployment.Spec.Template.Spec.Containers[0].Resources,
					VolumeMounts:    deployment.Spec.Template.Spec.Containers[0].VolumeMounts,
					Lifecycle:       deployment.Spec.Template.Spec.Containers[0].Lifecycle,
					Ports:           deployment.Spec.Template.Spec.Containers[0].Ports,
				},
			},
			DNSConfig:          deployment.Spec.Template.Spec.DNSConfig,
			SecurityContext:    deployment.Spec.Template.Spec.SecurityContext,
			ServiceAccountName: deployment.Spec.Template.Spec.ServiceAccountName,
			Volumes:            deployment.Spec.Template.Spec.Volumes,
		},
	}
	return pod
}

// getDeploymentList 获取所有的deployment
func getDeploymentList(deploymentInformer informerAppsV1.DeploymentInformer) (deploymentList []map[string]string, err error) {
	deploymentsZheliDev, err := deploymentInformer.Lister().Deployments("zheli-dev").List(labels.Everything())
	deploymentsZheliTest, err := deploymentInformer.Lister().Deployments("zheli-test").List(labels.Everything())
	if err != nil {
		return
	}
	deploymentList = make([]map[string]string, 0, len(deploymentsZheliDev)+len(deploymentsZheliTest))
	for _, v := range deploymentsZheliDev {
		deploymentList = append(deploymentList, map[string]string{"label": v.Name, "name": v.Name})
	}
	for _, v := range deploymentsZheliTest {
		deploymentList = append(deploymentList, map[string]string{"label": v.Name, "name": v.Name})
	}
	return
}

// getDeploymentController 获取deploymentList的控制器
func getDeploymentController(c *gin.Context) {
	deploymentList, err := getDeploymentList(DeploymentInformer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deployments": deploymentList})
	return
}

// deleteExpiredDebugPod 删除过期的debug pod
func deleteExpiredDebugPod() {
	// appLabel debug: true
	appLabel := map[string]string{
		"debug": "true",
	}
	// lister
	pods, err := PodInformer.Lister().Pods("").List(labels.SelectorFromSet(appLabel))
	if err != nil {
		log.Println("deleteExpiredDebugPod List error", err)
		return
	}
	for _, v := range pods {
		if v.Namespace != "zheli-dev" && v.Namespace != "zheli-test" {
			continue
		}
		// 15分钟
		if time.Now().Unix()-v.CreationTimestamp.Unix() > 900 {
			err = Clientset.CoreV1().Pods(v.Namespace).Delete(context.Background(), v.Name, metav1.DeleteOptions{})
			if err != nil {
				log.Println("deleteExpiredDebugPod Delete error", err)
				continue
			}
		}
	}
}

// deleteExpiredDebugPodLoop 定时删除过期的debug pod
func deleteExpiredDebugPodLoop() {
	for {
		deleteExpiredDebugPod()
		time.Sleep(time.Second * 60)
	}
}
