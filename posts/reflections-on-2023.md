---
title: "A Year in the Life of a Junior Developer"
date: 2023-12-29
author: "Harry Day"
tags: ["internship", "career", "year-in-review"]
slug: "2023-review"
description: "A reflection on the year of 2023, from the perspective of a junior developer"
---
## Introduction 
A lot has changed in 2023 in my life. I have finished my second year of university, started my industrial placement year at UBS, moved to London from Manchester and much more.
This article is a summary of the things I have learned and experienced as well as my outlook for 2024.

## Major 2023 Milestones

### Securing my placement and finishing year 2 of university
The start of 2023 saw me submitting my final applications for year-long placements as part of my degree programme. You can read more about this process in my other article: [How I Landed an Internship as a Software Engineer at a Leading Financial Institution](https://harryday.xyz/posts/how-to-get-an-internship/).
After the placement was secured, I focused back on university as well as finding a place to live in London, as I would be moving in June. This was alongside completing my second year of university, where I had a group project in Java and some other individual modules. 

### Moving to London
Soon enough, June rolled around and it was time to pack my life up and move it London. This process was smooth enough and a few weeks later I was standing outside of the building that would become my office for the next year.
The first month was filled with training, onboarding and getting to know the team as well as the projects that they are working on. I wanted to make an impact and proposed building a new library which would enable my team, and other teams within the function, to be able to access some of our services. This project occupied most of my time for the first month and was quite successful with mine and multiple other teams adopting the library within their services and applications with ease.

### Forming a new team
As I started to grow more comfortable within my role, I picked up more varied tasks and spent a lot of time working on our CI/CD pipelines. This was something that interested me and I strived to improve our pipelines and make them more resilient and useful. This work was mirroring other DevOps and platform-related tasks being completed by a few other members across the department and this led to us forming a new team focusing on enabling other application and services teams to be able to work at their maximum potential.
This has been a great experience being involved in the formation of a new pod, as it has allowed me to be a part of conversations shaping the future of how this team will operate and support other developers. 

Moving to this new smaller team placed me in closer, more consistent contact with a couple of far more experienced senior developers. Whilst they didn't agree on much, the different perspectives (from two different careers) were and still are extremely useful sources of knowledge. I have spent much of my time in this new team trying to soak up as much knowledge and experience from these key figures and I am forever grateful for having this opportunity. 

### Advent of Code
Finishing off the year strong, I decided to participate in [AoC 2023](https://adventofcode.com). The main intention behind this was to learn a new language: Go. Whilst I have not yet finished the whole challenge (the last 4 days were too busy in the lead-up to Christmas), I do think it has been useful. My proficiency in GoLang has skyrocketed from absolutely nothing at the end of November, to being confident in the language within a month. I will talk more about Go and my experience later in this post. 
I intend to finish off the remaining challenges at some point in the new year, however, I am happy with what I have accomplished so far (the maths and algorithms are starting to get a bit ahead of me now).

## My Developer Experience

Over the last year, I have grown more than I have in any previous year as a software engineer and developer. I think this is in no small part due to my 9-5 role as part of my internship. This has led to me spending far more time writing code than I have in any previous year where I have been busy with school/university.

> At the time of writing (29th December 2023) I have committed 1936 times. 

### Languages
This year I have learnt a few new languages as well as improved at ones I already knew. 
At university, our group project was written in Java, however, we also gained exposure to plenty of other languages, covering the basics of C/C++, Rust, Haskell, Java and Python. Whilst some of these I knew already, gaining some experience in Rust and Haskell as well as more time to try out C++ was useful. I think that Rust is something I would like to explore more of (especially after watching ThePrimeagen) as this was only covered very briefly at the end of my second year of university. 

At work, we write almost everything in C# and DotNet. Overall I do not mind this language, I feel like I can accomplish everything I need to without being limited. DotNet can be a daunting framework to come up against, but with the help of other more experienced developers around me, I have navigated learning the basics well and as always learnt plenty about how different people believe things should be done.

Advent of Code, as mentioned earlier, was something I used to help provide some challenges as I began to learn GoLang and I think this was rather successful. There are of course still more complex parts of the language that I haven't yet come across needing but hopefully, these will all be explored in time. I have many thoughts on GoLang, but ultimately I do like the language. I enjoy the explicitness around errors and handling them compared to throwing exceptions (I think that Swift does a good job with this style of error handling). I think that sometimes the STL can let Go down, for example, a set is a common structure I needed during AoC and I ended up having to use a dictionary which worked but wasn't as nice syntactically.

```golang
var someSet = make(map[string]struct{})

someSet["someKey"] = struct{}{}

_, prs := someSet["someKey"]
if prs {
    // someKey is in the set
}
```

I think that for 2024, my focus will be split across a few main areas:

1. Continuing my development using Go and Swift as these are both languages I enjoy using and think are quite capable.
2. Learn the basics of Rust and OCaml, I want to build a few projects using each to gain some more experience with them
3. Continue to develop my C# and DotNet knowledge whilst on my internship

### Cloud and Kubernetes
Working inside a large firm has also provided me with exposure to complex cloud environments that I otherwise would not be able to. In particular, one of the senior developers I work with is a Kubernetes nutcase and has taught me plenty about how it works and best practices. This has been a useful experience and has led me to be more mindful about how an application or service I am writing would be deployed to some sort of containerised environment.

I do however also see that some teams fall into the trap of putting absolutely everything into an AKS cluster and spending far more money than they need to just to deploy one simple service. With great power comes great responsibility (and bills) and I do see K8 being overused when a simpler solution would also fit the requirements just fine. 

In the new year, I want to focus more on how we manage and deploy our cloud environment at work. In particular, my team is leading the charge with the adoption of Terraform in the firm so this shall hopefully provide an exciting opportunity to learn and develop some documentation and tooling for other teams to use. 

### Tooling
For the past few years, my editor of choice has been Atom. Once this was sunsetted I switched to VSCode but was never quite satisfied with it. I picked up the JetBrains suite of tools as I am eligible for their student offerings and liked the built-in features dedicated to the respective language. IntelliJ IDEA got me comfortably through my Java work at uni (despite the professors' suggestions to use Eclipse🤮). 

Entering the back end of the year, I was red-pilled into switching from the traditional editor experience and installing Vim plugins. I committed to this across all my editors both at home and at work. 
I am slowly getting faster and faster, whilst discovering new motions I didn't know about. Whilst the initial impact slowed me down drastically, I do think I am writing faster than I have before. It is just a matter of slowly unlearning bad habits (like touching the mouse to move the cursor). 

This was put into effect around October and I think that Vim motions plugins for Visual Studio and IntelliJ was the right step, however, there was still the golden goose that I was ultimately chasing. ✨NeoVim✨

So that's why at the start of December I installed NeoVim and started to use it. My strategy was not to go cold turkey on normal editors (this is impossible anyway as I am unable to use NeoVim at work and therefore have to stick with Visual Studio there), but I slowly started to transition. Advent of Code was still mostly written in GoLand by JetBrains as I have still not figured out how I want debugging to work in Vim. 

Just yesterday, I rewrote my NeoVim config from scratch by hand. Originally I started with just using Kickstart (and boy am I glad I did that and not go down the NVChad route) however I wasn't 100% happy with the keymaps and plugins. Rather than trying to Frankenstein a custom config from Kickstart, I wrote my own.

It is based heavily on the things I liked from Kickstart, however includes more personalised keymaps, less bloat with some things I don't need, and of course [Harpoon](https://github.com/ThePrimeagen/harpoon).

I am sure my [config](https://github.com/harrydayexe/dotfiles/tree/main/nvim) will evolve as my needs change, but for now, I am happy with it. My end goal is to cut my reliance on other external editors and write everything in NeoVim. For now, this is not possible at work due to limitations out of my control, however, at home and when I am back in university in September, this is certainly my goal.

## Conclusion
2023 was a great year for me as a developer and as a human. I have achieved more than I thought I could and I am looking to do the same again in 2024. 
A renewed focus on my health, sleep, and fitness, alongside a renewed passion for software engineering, shall hopefully drive me to my goals. 

If you are interested in reading more, I upload a blog post semi-frequently (whenever something interesting comes up to write about) however I am aiming to do this more. My other posts can be found [here](https://harryday.xyz/posts/).

Check me out in other places on the interwebs:
- X [@harrydayexe](https://twitter.com/harrydayexe)
- GitHub [@harrydayexe](https://github.com/harrydayexe)
- LinkedIn [/harrydayexe](https://www.linkedin.com/in/harrydayexe)
- (Are you seeing the trend yet)
