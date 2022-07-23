# Dcard-Like-System

## 1.Create a Dcard-like System
* Providing users a platform to communicate.

<br>

## 2. Requirements and Goals
### Functional Requirements
> Member (Required)

> Posts (Required)
>> * Post upload
>> * Attachment upload (e.g. images, documents)
>> * Comment
>> * DM (Bonus)

> Draw Card (Required)
>> * Target must be opposite sex
### Non-Functional Requirements
> * 

<br>

## 3. Capacity Estimation and Consstraints

<br>

## 4. System APIs
### Upload post
    func postCreate(memberID, title, content string, fileLink []string, likes int, postDate datetime)(statusCode int)
### Get all posts
    func postsGet(page, perPage int)(postID, title, content string, likes int)
### Get single post
    func postGet(postID string)(memberID, title, content string, fileLind []string, likes int, postDate datetime)
* Post
> Function Intro
>> * Post Upload
>> * Comment
>> * Like

* Draw Cards
> Function Intro
>> * Matching
>> * Send Request
>> * Matching Info

* Membership
> Function Intro
>> * Login/Logout
>> * Register
>> * Info CRUD
  
  

![1](https://user-images.githubusercontent.com/71340325/174738452-cc0a5f58-f65d-4f6e-9938-f07e0988e938.jpg)
![2](https://user-images.githubusercontent.com/71340325/174738529-0fc1498e-5c83-40dc-af65-210898618c53.jpg)
![3](https://user-images.githubusercontent.com/71340325/174738539-35ebfc7f-51ce-4f54-afe0-74c97e44ef9f.jpg)
![4](https://user-images.githubusercontent.com/71340325/174738558-b8e2132a-0ef8-4b00-9fc4-b8c1aebafc55.jpg)
![5](https://user-images.githubusercontent.com/71340325/174738571-3f23bc12-baee-4da5-8f80-97546ce45c4b.jpg)
![6](https://user-images.githubusercontent.com/71340325/174738578-95d0a678-064e-4cf4-9179-85bae4241a01.jpg)
![7](https://user-images.githubusercontent.com/71340325/174738699-2236cff5-e302-4b41-8b7b-34853d628422.jpg)
![Login](https://user-images.githubusercontent.com/71340325/180591814-94cbfb4a-ed00-4128-b6f2-9011fac31b39.jpg)
![9](https://user-images.githubusercontent.com/71340325/174738730-7e788e47-519b-4613-8480-e2698836bb1e.jpg)
![10](https://user-images.githubusercontent.com/71340325/174738736-c1eb82ef-ea1f-44c2-a3ff-f45603533a8d.jpg)
![11](https://user-images.githubusercontent.com/71340325/174738753-d8a7dfab-456b-4c7f-8634-51dfb866a6b2.jpg)
![12](https://user-images.githubusercontent.com/71340325/174738759-19fecd4e-8025-4b38-b27e-5887d370bb72.jpg)
![13](https://user-images.githubusercontent.com/71340325/174738775-8627678e-33e1-487e-b03b-411a5f62edb5.jpg)
![14](https://user-images.githubusercontent.com/71340325/174739242-610bc1b5-0b92-449b-931a-8562d70abfc0.jpg)
![15](https://user-images.githubusercontent.com/71340325/174739253-57db34f0-7e0d-4b5d-ac97-4d49d7d4da01.jpg)
![17](https://user-images.githubusercontent.com/71340325/174739283-f623734b-033c-49ec-8074-eea22a3d3182.jpg)
![18](https://user-images.githubusercontent.com/71340325/174739296-e53a9a59-ff94-410e-94f0-a36580a64b19.jpg)
![19](https://user-images.githubusercontent.com/71340325/174739325-afe9a152-5156-4097-9be8-5a332980070b.jpg)
![20](https://user-images.githubusercontent.com/71340325/174739338-fe1cf083-a75b-4938-a0b9-75caf6ba6f62.jpg)
![21](https://user-images.githubusercontent.com/71340325/174739351-f2bd24f8-38dd-4310-9295-aee2ab8a7b8a.jpg)
<br>


## Commands

Enable API Gateway
> go run main.go apigateway

Enable Post Server
> go run main.go post

Enable Matching Server
> go run main.go matching

Enable Member Server
> go run main.go member
