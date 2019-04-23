# Kubernetes Deployment Example

## This exercise is all about Kubernetes deployment. Here we have go api service deployment step by step.

**1.** Create an api service using GoLang. Here we have main.go

**2.** Create a Dockerfile to dockerize the api service.

  ```
  FROM golang
  RUN mkdir app/ && cd app/
  WORKDIR app/
  COPY . .
  RUN go build main.go
  CMD ["./main"]
  EXPOSE 8085
  ```
  
  **3.** Buid an image using above docker file-
  
          
          docker build -t <imagename:tagname> .
          
  
  **4.** Now before push this image to docker hub, we will do the tagging of image.

         
         docker tag <image-id> <reponame:tagname>
         
  **4.** Push the image to your docker hub repository that you built using-
  
         
         docker push <reponame:tagname>
         
  
  **NOTE:** Before pushing image to docker hub, you must login to docker using- ```docker login``` command
  
  **5.** In this step we create 1 YAML file with name deployment.yaml
  
        
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: api-deployment
        spec:
          replicas: 2
          selector:
            matchLabels:
              app: api-deployment
          template:
            metadata:
              labels:
                app: api-deployment
            spec:
              containers:
              - name: api-deployment
                image: sannidocker/myfirstrepository:api-image-v1.0
                ports:
                 - containerPort: 80
             
             
  **6.** Now create deployment using deployment.yaml. This will create a deployment with 2 replica pods.
  
         
         kubectl create -f deployment.yaml
         
         
  **7.** To check the deployment and pods has created or not, we can use the below commands
  
         
         $ kubectl get deployments
           NAME             READY   UP-TO-DATE   AVAILABLE   AGE
           api-deployment   2/2     2            2           26m
           
         $ kubectl get pods
           NAME                              READY   STATUS    RESTARTS   AGE
           api-deployment-5f9c9cfb4d-7pdd7   1/1     Running   0          25m
           api-deployment-5f9c9cfb4d-nnrmp   1/1     Running   0          25m
         
         
  ### Scale-up and Scale-down
  
   Scale-up and scale-down can happen in two ways:
    
   a. ```kubectl edit deployments <deployment-name>```
   
   - Here deployment-name is same as given in deployment.yaml file
        
   b. ``` kubectl scale <deployment-name> --replicas=<number-of-replica> ```
   
   - Here deployment-name we can get using ```kubectl get all```
        
        ```
        $ kubectl get all
        NAME                                  READY   STATUS    RESTARTS   AGE
        pod/api-deployment-869cb9548c-bxssz   1/1     Running   0          29m
        pod/api-deployment-869cb9548c-nvz82   1/1     Running   0          29m

        NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
        service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   11d

        NAME                             READY   UP-TO-DATE   AVAILABLE   AGE 
        deployment.apps/api-deployment   2/2     2            2           29m

        NAME                                        DESIRED   CURRENT   READY   AGE
        replicaset.apps/api-deployment-869cb9548c   2         2         2       29m
        ```
