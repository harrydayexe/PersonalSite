---
title: "My First Hackathon"
date: 2023-10-24
author: "Harry Day"
tags: ["hackathon"]
slug: "2023-hackathon"
description: "Describing my first hackathon and the project I created."
---
Last week I had the opportunity to take part in a global hackathon at the bank where I work. This event took place over two days, and I entered as part of a team of junior talent, made up of interns and grads. 

The brief for the hackathon was rather simple. There were a few categories you could build a project under, or a “BYOI” (bring your own idea) category which was a catch-all for anything else. Building something to improve the bank. 

## The Problem
We have an internal tool that maintains authorisation and access rights to practically everything in the bank. Need write access to a particular GitLab group? You need to submit the right request. Need to be able to send external emails? You need to submit the right request. The internal tool has a rather unfriendly (to put it politely) interface and is very confusing for new joiners to understand and use. 
As a team of relatively new joiners, we understood the pain of this system well and were determined to improve it. 

## The Solution
Our team’s plan was simple. We would create an AI chatbot that you could ask a question to, and it would provide the relevant permission code to request. For example, telling the bot “I want to edit and maintain X GitLab repo” would return a response along the lines of “To access this repo, you should request X permission with Y authorisation code”. 

In a production version of this application, this data could be scraped from existing internal documentation pages, however, for the proof of concept that we needed to deliver, we decided to pull out a couple of examples of these requests and provide them directly to the AI. 

## The Design
As part of the hackathon, the bank made available an internal AI platform which utilised OpenAI’s tools. We used a stateless system, maintaining a queue of system messages and conversation history in the backend, with this being passed in each API request. This means that no potentially confidential information was stored in OpenAI. 

These API requests were driven from a NodeJS server. This server monitored web sockets for incoming user messages, which would then be formatted into a request to the OpenAI API. Once a response was received, this was passed back to the client through the web socket.

On the client side, we used React to build a simple chatbot interface that could monitor a web socket for incoming messages, and display them on the screen. 

## Implementation
On the first day of the hackathon, we sat down and quickly planned out what each person would be doing. I was allocated to working on the server side NodeJS work. The first day went well, we started programming around 11 am and by about 8 pm I had finished the server. In parallel, the front-end team developed an initial react app up until they called it a day around 9 p.m. 

Coming into day 2, our code freeze was 1 pm. We started full steam ahead at 9 a.m. and realised we still had lots of work ahead of us. Plagued by npm issues on multiple laptops, we eventually got both the client and server working on the same machine at the same time. Once communication was working via the web socket, it was time to merge the code into the main branch and get ready for our presentation. 

## Presentation
We had a quick lunch and then headed into the auditorium for the presentations of the projects created by the London-based teams. Our turn was about halfway through and we went for a Dragon’s Den style pitch to spice things up. An explanation of the problem we were trying to solve, a quick demo and a few jokes later and we were headed off the stage to a round of applause. Speaking to some members of the audience, they seemed to like our presentation but there was tough competition from other teams with very useful projects so we weren’t overly confident.

## The Results
In the end, our team was not shortlisted for the finals as the competition was tough. We certainly learned a lot over the 2 days and it was a great experience overall. As a team comprised of only junior talent, I would say that we held our own and could be proud of the work we produced. It is something I would try again.
