# exaBase Studio Take-home Exercise

Thank you for your interest in joining the exaBase Studio team!

Going through this exercise will not only help us evaluate your approach to solving engineering tasks, but it will also give you a basic understanding of the product that we're building.

## Instructions
* Once you're done with the exercise, please send your solution (e.g. as a zip file) to the main point of contact for your recruitment process.
* Feel free to use any programming language you're familiar with for solving this exercise. However, don't assume that the reviewer is also familiar with the language, and aim for the code to be readable and easy to understand.
* If you notice any "unclear" requirements in the exercise, feel free to mention this in your solution and talk about options you're considering and assumptions you're making.
* There are several "stages" in the exercise, with several workflow examples for stages 2-3; please describe how we can review and run each stage's solutions.
* The exercise should typically take around 3 hours. If you spend significantly more time on it, please let us know.

## Context
exaBase Studio is a low-code graphical environment for developing AI-enabled workflows, and providing facilities to deploy them automatically, ready for production use. Developing and deploying such a workflow involves several layers:
1. Canvas - a graphical representation of the workflow, created in the web UI of exaBase Studio. It typically contains definitions of input/output endpoints, visual representation of data flow between elements, and references to dockerized ML models that will interpret input data and provide useful results as output.
2. Blueprint - a yaml representation of the workflow, generated from the Canvas above, to be passed as an argument to Constructor.
3. Constructor - the main orchestrating component, taking the Blueprint from the layer above and deploying it to a Kubernetes cluster as an executable Circuit.
4. Circuit - a "live", working workflow, as defined in the Canvas/Blueprint, with accessible API endpoints for inputs and outputs, ready to ingest data and kick off processing in the ML models it includes.

The aim of this exercise is to build a vastly simplified version of exaBase Studio, touching on the Blueprint and Constructor layers, and allowing for simple execution of logic, similar to how it happens in the Circuit.

Let's call it __miniStudio__!

## Exercise

### Stage 1 - warm-up

This stage will be simple, but setting the foundation for the rest of the exercise. The goal - __add two numbers!__ ...but with a couple of twists.

exaBase Studio uses API endpoints to post data into the workflow and receive end results. For the purpose of this simplified exercise, we can __use plain text files to simulate these endoints__. Create the following files:

For inputs:
* `input1`, in the beginning containing just one number, e.g. `2`
* `input2`, containing another number, e.g. `3`

For output:
* `output1`, empty in the beginning

To simulate a Circuit running and listening for new data, we can use __a simple infinite loop__ reading from the input files.

#### Explanatory flowchart
_Flowcharts in this document are rendered using Mermaid Live Editor; just follow the link below to see the chart._  
[Stage 1 flowchart](https://mermaid.live/view#pako:eNo1jTELgzAQhf9KuMmCDnF06FSEQqFgR-NwNWcVTCLpBSnif2_U9qb3Pj7eLdA6TVBAN7q57dGzuFXKinhXmQx2CixPIsvOopR1jVoLnp2wwTzJv5vmZ-aHmf_NA5dyr3eZuMD70MYhBUPe4KDj12UjCrgnQwqKGDV1GEZWoOwaVQzsHh_bQsE-UAph0sh0GfDl0Rxw_QLGXD00)

#### Acceptance criteria

* Have the infinite loop of your program read both input files, `input1` and `input2`, parse numbers contained in them, add them together, and write the result to the output file, `output1`
  - The initial result would be that `output1` should be filled out with `5`.
  - While the infinite loop keeps running, if you post another value to one of the input files, the output value should change; e.g. after running `echo 10 > input2`, the `output1` file should be updated with `13`.

### Stage 2 - different functions

Now we want to make our miniStudio more flexible, making it possible to "plug in" other functions for processing inputs.

#### Explanatory flowchart
[Stage 2 flowchart](https://mermaid.live/view#pako:eNo1jcsKgzAQRX8lzMqCLuLSRVdFKBQK7dK4SJOogTwkzFCK-O-N2s7q3sPhzgIqagMNDC6-1SQTsttDBJbvygsbZkJ-YlV1Zi3vutnROMqXM2ygoNDG0Pc_uT7k-i8fuOV7vfMiEu5bG4cSvEleWp0fLxsRgJPxRkCTozaDJIcCRFizKgnj8xMUNJjIlECzlmguVo5J-gOuX887Pq4)

#### Acceptance criteria

* Allow declarative definition of your workflow's "Blueprint", taking a function as an argument.
* Prepare a "function library" with "pluggable" functions. For now these functions would all take 2 arguments (as previous `input1` and `input2`) and provide one result (later to be written to `output1`).
  - Extract the previous logic of adding two numbers into its own function, e.g. `addNumbers`.
  - Prepare another function that will multiply two numbers, e.g. `multiplyNumbers`.
  - Prepare yet another function that takes two strings and concatenates them, e.g. `concatenateStrings`.
* Keep the inifinite loop approach from Stage 1 to simulate a running Circuit listening for new data.

##### Examples
* When initializing your "Circuit" with a "Blueprint" containing the __`addNumbers`__ function:
  - Put `1` in `input1`, `100` in `input2` - the result of `101` should appear in `output1`.
  - Put `1000` in `input2` - the result of `1001` should appear in `output1`.
  - Put `50` in `input1` - the result of `1050` should appear in `output1`.
* When initializing with __`multiplyNumbers`__ function:
  - Put `4` in `input1`, `8` in `input2` - the result of `32` should appear in `output1`.
  - Put `1000` in `input2` - the result of `4000` should appear in `output1`.
  - Put `50` in `input1` - the result of `50000` should appear in `output1`.
* When initializing with __`concatenateStrings`__ function:
  - Put `4` in `input1`, `8` in `input2` - the result of `48` should appear in `output1`.
  - Put `melon` in `input2` - the result of `4melon` should appear in `output1`.
  - Put `water` in `input1` - the result of `watermelon` should appear in `output1`.

### Stage 3 - final form
So far we've been working with 2 inputs, 1 function, and 1 output. The real exaBase Studio allows specifying more complicated data flows, operating on multiple inputs and outputs, with multiple chained functional components.

An additional element that allows more complicated dependencies between parts of the workflow is what we call an "entity". It also allows for some data persistence. In this Stage of miniStudio we will be using entities as "stores" for data going in and out through endpoints, as well as for passing data between chained functions.

#### Acceptance criteria
* "State" stored in entities only needs to be persisted while the program is running - no disk/database persistence necessary.
* Implement a flexible solution allowing declarative definition of workflows presented in the flowcharts below; keep the "Circuit" (execution engine) and workflow definition as separate concepts.

##### Examples to fulfill
* [Add and accumulate incoming numbers](https://mermaid.live/view#pako:eNo9T02rwkAM_CshJwWlbI89eLKCIAi-Y7eH0E11od0t2-x7iPjf39oPcwiZzDCZvLDxhrHAtvN_zYOCwOWmHaQ6q411QxS1hf3-AKWqMnZi5Qkqq2dJqSbqpKqKjIFf6iKP4FtYhUDOrCCfwCg-MAQeYydg3ZesF8vTbFnmy4V8ubDCb4g8qyfqqjY-SsoJaqvdLPt03GHPoSdr0nevz0ajPLhnjUUaDbeUImjU7p2kFMX_PF2DhYTIO4yDIeGjpXugfl6-_wHp9lu8)
* [Multiply multiple numbers](https://mermaid.live/view#pako:eNpVz8GKwjAQBuBXCXNSUCRJTz3syQqCILjHxkO2nWqgSUucsBTx3Te2MbA5ZWY-fmae0AwtQgldP_w2d-2JnS7KsfiOfGXcGIiv2Xb7xSpe79CRoYnx3TURsRCRiMhEZCIXIhORmcgPqfg8OvC6tqEnM_YTc8H-oH9cP0Qkkkr5rzwsAVWRs4ucXcyjM18NgeZj3n3YgEVvtWnj5c93RwHd0aKCMn5b7HRcRIFyr0h1oOF7cg2U5ANuIIytJtwbffPaQtnp_oGvP5Y0YH8)
* [Concatenate strings and keep track of longest result](https://mermaid.live/view#pako:eNqNkUFrxCAQhf_K4GkXNiya5JJDT02gUCi0x2QPYkxWiBp0pCzL_vfaNUlhaaGCoPO9efPQKxG2l6Qiw2Q_xZk7hNf3zkBcL3SnzByQ7iHLnqCm7VEaVHgBejwtEpYkbJGwTcJWSU3vqKFtK6wRHKWJGzw6ZUZ_WlVsUaVrk5rqfPPLIQMRnItX-PFR1oCTPky4jcuTEWtbj9ZJmKwZpf-9CbyFgbs1RJNC1MU2tYhT_2FwfDAoN4PyTwNUOha5nh-Sv9GdDRjfFOh-AUUCbAVsBWUC-QryOyAHoqXTXPXxW6_flY7gWWrZkSoeeznwGL0jnblFKQ9oPy5GkApdkAcS5j6GfFZ8dFyn4u0Lq2-qqA)

### Stage 4 - user-facing documentation
You're done with all the implementation stages, congratulations! The last request we have for you is to prepare a bit of documentation. Describe what someone new to miniStudio needs to know to be able to:
* run the program,
* define their own workflows,
* define new functions to use in their workflows.

## Final thoughts
While we've been operating on simple functions for our simplified miniStudio, the real exaBase Studio easily works with powerful ML models and e.g. GPT-like engines to provide our customers tools to build their own AI-powered tools and processes.

We hope you enjoyed this journey of building miniStudio, and that you now have a slightly better understanding of the product that our team is building. We are looking forward to seeing your approach to implementing this system!
