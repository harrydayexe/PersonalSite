---
title: "Building GoBlog"
date: 2026-06-22
author: "Harry Day"
tags: ["ai", "golang", "open-source"]
slug: "building-goblog"
description: "Setting foundational architecture to keep AI in check."
---
At the start of 2026 I set out to rebuild my website from scratch. The old version was built by hand years ago using a combination of HTML, raw CSS, blood, sweat, and many, *many* tears. 
This was all hosted on a shared hosting platform, by manually uploading the files from my laptop.

While teenager me was happy with my creation, adult me knew there was a better way.
After all, I have built many websites for others, using the latest and greatest frameworks; but I have never treated myself to such an upgrade.

## Requirements For My Blog
The first thing any responsible engineer should do, especially with personal projects that actually need to be finished, is to set out a list of requirements. This would give me a clear end goal to actually hit publish and go live with the new site, rather than dooming it to an eternity of living in my dev folder.

1. **Hosted on a VPS**
   This was a no-brainer for me. I already had a spare VPS on hand and it meant I controlled every inch of the software stack; no more messing around in cPanel trying to figure out how to upload my site. I'm a software engineer, not a web administrator.
2. **Easy blog generation**
   With a new focus on wanting to write blog content (thanks for reading), I wanted a way to easily take new articles and get them uploaded as seamlessly as possible. I had experimented with projects such as John Sundell's [Publish](https://github.com/JohnSundell/Publish) in the past, but I never quite committed to it.
3. **Markdown articles**
   I use [Obsidian](https://obsidian.md) as my note-vault of choice, and markdown here is king. I want all my articles to be described in a standard format that can be adapted anywhere.
4. **Customisablity as standard**
   I want the ability to easily change the entire look and feel of my website, or just add a small feature, without requiring a huge rewrite, and without changing any of the blog content itself.

Based on these requirements I decided to build my own open-source framework for blog generation, which could be tweaked and perfected over time: [GoBlog](https://github.com/harrydayexe/GoBlog). Ultimately the decision to build in Go came down to familiarity, it has been my language of choice for the past couple of years, and excels in the web domain.

> I didn't spend much time exploring the existing options already available that may have fulfilled these requirements. Whilst modernising my website is the ultimate goal, I also wanted to use this as an opportunity to build something new, and explore how I can utilise AI without giving up control.
## An Aside on AI
I will preface the rest of this article by saying that in many ways I am a traditionalist when it comes to software engineering, however I am not blind to the direction the field is moving in.

I was in my first year of university when ChatGPT went mainstream. I watched the way that those around me used it as crutch. To me, the point of studying computer science at university is not to learn all the latest and greatest technologies and frameworks, the world moves far too rapidly for a syllabus to keep up. It is instead to learn and understand the basics that underpin modern computing. 

Watching others around me use ChatGPT to complete their assignments, I felt strongly about not letting a tool replace my core knowledge, betting on the fact that LLMs would serve me better if I actually understood the core stuff. This is the same ethos I carry with me now, and is something I argue regularly with the "true believers".

## Vibe-Coding: An Initial Exploration
My first attempt at building GoBlog was to go all-in on AI. I wanted to see how far I could take vibe coding a framework before it all started to fall apart. 

Undoubtably the progress looked good on paper. Claude confidently produced code towards my goals with perfect test coverage and reassuring explanations of what it was doing. However, when I actually looked at the codebase, I struggled to understand the logic of why it was structured in the way it was. Some of this was certainly down to the way I was using AI, I hadn't used the available tools like CLAUDE.md files, skills and agents correctly, but this was a learning process after all.

That's not to say I haven't had success with AI, I used it extensively at [Novlr](https://www.novlr.org), during my time with the company, without facing the same kind of issues. So what then is the problem? Stronger guardrails and better use of the available tools? Sure. But my bigger belief is that it comes back to the old-adage: sh\*t in sh\*t out.

I decided that if I wanted to use AI effectively, I would need to properly configure Claude Code, and give it a basic foundation of what I wanted my software to look like. 

## GoBlog v2
I decided to scrap everything I had built so far for GoBlog and start from scratch. I first used the Claude desktop app to work together on drafting a general outline of how GoBlog would work. I have found that my preferred workflow for initial design/planning is to use a model like Opus to draft a first revision based on all the information I can throw at it. Often times it will explore the web and throw new ideas back at me, which will then inform the next iteration. 

> In my experience, Anthropic's models are much better at not aiming to please, they are more willing to push back on objectively bad ideas, where models from Google and OpenAI are more likely to just agree and force the idea to work, often resulting in a messier, worse final outcome.

The back and forth here set out a general layout and some goals for GoBlog itself. I knew that I wanted to build a CLI tool which could both generate and serve blog content, as well as an accompanying go module which could be woven into other web apps for more customised and controlled usage.

Ultimately to build GoBlog, I needed a few things:
- A parser to take in content and extract the information it needs
- A generator which takes the output of the parser, and inserts it into templates
- A server to respond to web requests and send back the generated pages

These 3 key sections formed the top-level of my API surface, but it is by no means an exhaustive list. 
> You may be wondering why this is being called v2, and not just a second attempt at v1. Well due to already registering my first attempt of GoBlog with [pkg.go.dev](https://pkg.go.dev), I faced many issues with trying to overwrite a new release of the v2 code to the same release versions as v1. Due to this I made the decision that GoBlog will start at `/v2`. Honestly I should have just picked a new name, but I didn't think it would bother me as much as it has ended up doing.
### General Project Setup
One of the earliest things implemented, and mostly by hand, was my GitHub workflows. These were mostly written by hand for two reasons:
- If you are relying on AI to write your code, do not rely on AI to write the gates that prevent bad code from merging into main
- AI *sucks* at ensuring it is using the most up-to-date versions of GitHub actions

Because of these two reasons I decided to setup the workflows myself, but mostly because I was fresh off the heels of being burned in v1 and I wasn't ready to give up control yet.

With the pipelines in place to build, test and release GoBlog, I was ready to continue with the general setup of any new project: a `.gitignore`, [just](https://github.com/casey/just) for running commands, [goreleaser](https://goreleaser.com) for releases and, of course, a `.claude` directory.

### The AI Workflow
At this point, I was ready to get into the weeds and actually make some progress. I set up some general submodule directory structure to keep Claude on track, and started with the parser implementation first. 

> I will spare the details of how everything is implemented, that's not my intended focus of this article; if you're interested the code is open-source under the [MPL-2.0 License](https://www.mozilla.org/en-GB/MPL/2.0/).

The general workflow in the beginning of GoBlog looked something like the following:
1. For each big new submodule, I would first chat with Claude to draft a general design guide and overview for how I wanted something implemented.
	1. This covers everything from general features, acceptance criteria, suggested design etc.
	2. This guide was exported as a final markdown document and added to my Obsidian notes. Parts of it would be copied into Claude Code when needed.
2. For each new feature implemented in the codebase, I would follow the basic Plan-Implement loop.
	1. The plan would be done in plan mode using Opus, going through a refinement process until I was happy with it
	2. Once the plan was finalised, it would be handed off to Sonnet and any sub-agents to implement
3. Post implementation, I would sanity-check the work in a PR, often going through some smaller iterations without plan mode to fix any issues I have spotted
4. Once I was happy with a feature, this would be committed to main and the cycle would continue

> Git Worktrees are an under-valued feature of Git which make concurrent work on non-dependent tasks by AI possible. While one agent is planning, another can be implementing while you are prompting a third. I won't go into too much detail on worktrees in this article but the [docs by Anthropic](https://code.claude.com/docs/en/worktrees) are quite comprehensive.

In general, this process worked well. Focusing on one particular feature at a time with clearer and more well-defined guidelines on where it should be placed, how it should be implemented, etc made the resulting codebase a lot cleaner and easier to navigate. 

Personally, I think that codebases *should* be opinionated. Opinions can range from code location and module responsibilities, to code formatting, to overarching design decisions. Knowing when to say "this is not the responsibility, or priority, of this project" is important.

## The First User
I built GoBlog for myself. I hope that it may be useful for others, but ultimately this is not a business, it is completely open source and I am not looking to make money from it. 

Once I was happy with the MVP of GoBlog, I decided to start using it in my own website. Due to the blog being only one part of my website, using the API to embed it into my own application made the most sense.

While the overall aim of this section was to build a functional (and pretty) website, I also wanted to explore how well an LLM could use the documentation for GoBlog to utilise it. I think that it is important to test how easily others, who may be less technically skilled, can use your frameworks in a vibe-coded application themselves. 

During the development of this application, I noticed several gaps in the documentation, as well as discrepancies within it that caused LLMs to get confused. These gaps highlighted why it's always important to actually use the libraries you have built, and an LLM without the knowledge that you have of a project is a perfect test-dummy to see just how good your docs really are.

> I have noticed an interesting behaviour with Claude Code, where when it gets confused about the behaviour of a go module, it will attempt to find the cached source code on disk in order to read it. 
> Personally I prefer to not rely on knowledge of the internal workings of a library, in favour of relying purely on documented behaviour in the docs. I have gone through many iterations of `PreToolUse` hooks in order to block this. 

## Further Improvements
By this point, I am pretty happy with the functionality of GoBlog. There are still some features in the pipeline that I will work on when I have the time, but for now the vast majority of the work to build a usable product is finished. 

I have now adopted a more traditional approach for software engineering. New features are logged in GitHub issues with user stories and acceptance criteria. These are then processed in conversation with Claude Code to land on a general implementation plan.
In general, these kind of tickets do not require much intervention from me. Due to the strong foundations now set in the library, the implementation that I and an LLM decide upon are normally very closely aligned. 

I am considering building a tool which is able to work on these fleshed out specifications (which themselves are generated from the feature-request user stories) agentically. This would allow scheduling of the work to happen during the 5 hour usage limit windows where I am not actively working on something myself, increasing efficiency of the Claude subscription I pay for. 

> Ultimately I want to be able to wake-up to a set of merge requests opened by Claude, which cover all the features I drafted the night before.

## Areas to Improve
The one thing I did not design very clearly, and shows to this day, is the configuration setup of GoBlog. This was something I (stupidly) did not consider during the design phase, and as such has suffered the same consequences that v1 of GoBlog did. 

There is a config module, which contains *most* of the configuration for the public package. However it follows two main methods of setup at the moment and as such feels disjointed. I hope to rectify this in a v3 of the library, but, as this would be a breaking change, it is mostly on the back-burner for now.

> I do think that this proves the need for strong foundations and opinions when it comes to LLM code generation. Without this clear structure baked into the project, especially in the early stages where there is not much code to use as an example, it is very easy for the LLM to produce two different patterns for the same thing. 

## Closing Thoughts On The Use of AI
It is clear that AI is here to stay. That much is undeniable. However that does not mean that software engineers are no longer valuable, our skillsets useless, or that LLMs are a silver bullet for software. 

I saw [this post](https://www.linkedin.com/posts/georgejeffers_this-is-the-most-important-tweet-you-will-share-7463990257978466304-gsje) from a founder on LinkedIn recently where one line stood out to me:

> "Code is a solved problem. Engineering used to be the bottleneck for software companies, but AI has removed this."

It is my opinion that code is not a solved problem. If it was, people wouldn’t pay for software-as-a-service because they could just ask an AI to build it for them. 

My day job in a global bank is the perfect example. No amount of AI acceleration will ever replace the engineering calls where the foundational decisions are made. These are not the sort of decisions that affect a single function, or a single class, or a single micro-service. They are the decisions that lay the groundwork for every future development objective a division will have. 

Software engineering requires direction and accountability. Only a human can make those decisions. No regulator, CEO, or manager will accept the excuse "AI wrote it" when real-world, negative outcomes occur due to software. 

Lay the foundation correctly, and you can reap the benefits of AI far more easily. 

Thank you for reading.
