I tried completing the task as much as I could without referring to any external sources, or browsing the internet. But I must declare, that this task wasn't event properly clear to me. I do not understand the architecture of how this service (the one described in task) is meant to work.

Some questions I can't comprehend:

1. Where is the external translation service?
2. What output am I supposed to assume for unit testing the translator?
3. Where exactly is the translation error coming from? How am I suppoed to assume an error?

I've done as much of this task as I could within 30-40 minutes, given the portions of task that immediately made sense to me on the first reading.

I've implemented the following:

1. The cache: created a very simple map based local cache only for demonstration, which contains the entire query and it's result from translation.
2. It's **very** easy for me to deduplicate idential queries to reduce the load, but I'm not sure where the batch requests are coming from? Because I've been clearly asked **not** to modify the `main.go` file from where queries are being made. And I'm assuming that is the file you will use for testing. 

Am I supposed to write a new translator function for batch requests? And how will you test deplicate batch requests? 

### Finally

I thought a lot about the Paulcamper platform. As far as I can imagine your platform's architecture in my head, it should be a fairly CRUD platform, additionally with some machine learning maybe. So, a basic postgres database, with either REST or GraphQL, should be fine, coupled with an authentication engine and the obvious front-end.

I have experience in designing multiple such platforms, but I can't understand how this specific test relates to such a platform.

I understand that you might've simply wanted to test some basic coding skills, but since this task itself requires a lot more clarification, I just wanted to make my confusion clear over here.

Thanks!