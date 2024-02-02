package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"playground/logs"

	// "github.com/bytedance/sonic"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestK8s() {
	Init()
	deploymentName := "clotho-tuas-user-kr-pt"
	namespace := "spatio-pt"

	HPA(namespace, deploymentName)

}

func GetDeployment(nameSpace string, deploymentName string) {

	// Get deployment and pretty print info
	deployment, err := k8sClient.AppsV1().Deployments(nameSpace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		// pretty print error using json
		logs.Logger.Errorw("error:", err)
	} else {
		//logs.Logger.Infof("deployment: %+v", deployment)
	}

	runningReplicas := deployment.Status.Replicas
	readyReplica := deployment.Status.ReadyReplicas
	desiredReplica := deployment.Spec.Replicas

	logs.Logger.Infow("Deployments", "running replicas", runningReplicas, "ready replicas", readyReplica, "desired replicas", desiredReplica)

}

func GetHPA(nameSpace string, deployment string) {
	hpaClient := k8sClient.AutoscalingV2().HorizontalPodAutoscalers(nameSpace)
	hpa, err := hpaClient.Get(context.Background(), deployment, metav1.GetOptions{})
	if err != nil {
		logs.Logger.Errorw("error:", err)
	} else {
		logs.Logger.Infof("hpa: %+v", hpa)
	}

	result, err := json.MarshalIndent(hpa, "", "    ")
	// print HPA info with indent
	fmt.Printf(string(result))

}

func HPA(namaSpace string, deployment string) {
	hpaClient := k8sClient.AutoscalingV2().HorizontalPodAutoscalers(namaSpace)
	min := int32(2)
	max := int32(4)
	percentage := int32(70)
	hpa := &autoscalingv2.HorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment,
			Namespace: namaSpace,
		},
		Spec: autoscalingv2.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: autoscalingv2.CrossVersionObjectReference{
				Kind:       "Deployment", // or "ReplicaSet", "StatefulSet", etc.
				Name:       deployment,
				APIVersion: "apps/v1",
			},
			MinReplicas: &min,
			MaxReplicas: max,
			Metrics: []autoscalingv2.MetricSpec{
				{
					Type: autoscalingv2.ResourceMetricSourceType,
					Resource: &autoscalingv2.ResourceMetricSource{
						Name: "cpu",
						Target: autoscalingv2.MetricTarget{
							Type:               autoscalingv2.UtilizationMetricType,
							AverageUtilization: &percentage,
						},
					},
				},
				{
					Type: autoscalingv2.ResourceMetricSourceType,
					Resource: &autoscalingv2.ResourceMetricSource{
						Name: "memory",
						Target: autoscalingv2.MetricTarget{
							Type:               autoscalingv2.UtilizationMetricType,
							AverageUtilization: &percentage,
						},
					},
				},
			},
		},
	}
	result, err := hpaClient.Update(context.Background(), hpa, metav1.UpdateOptions{})
	if err != nil {
		logs.Logger.Errorw("error:", err)
	} else {
		logs.Logger.Infof("hpa: %+v", result)
	}
	resultStr, err := json.MarshalIndent(result, "", "    ")
	// print HPA info with indent
	fmt.Printf(string(resultStr))

	// existingHPA, err := hpaClient.Get(context.TODO(), "<hpa-name>", metav1.GetOptions{})
	// if err != nil {
	// 	_, err = hpaClient.Create(context.TODO(), hpa, metav1.CreateOptions{})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println("HPA created successfully.")
	// } else {
	// 	hpa.ObjectMeta.ResourceVersion = existingHPA.ResourceVersion
	// 	_, err = hpaClient.Update(context.TODO(), hpa, metav1.UpdateOptions{})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println("HPA updated successfully.")
	// }

}

func CreatePod() {

	// Get deployment and pretty print info
	_, err := k8sClient.CoreV1().Pods("spatio-pt").Create(context.Background(), nil, metav1.CreateOptions{})
	// pretty print error using json
	logs.Logger.Errorw("error:", err)

}
