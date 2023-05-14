# Behavior Questions
- [Behavior Questions](#behavior-questions)
  - [Personality](#personality)
  - [Team Work](#team-work)
  - [Career Plan](#career-plan)
  - [Jobs Experience](#jobs-experience)

## Personality
1. What do you think are your greatest strengths and weaknesses?？
2. If you have chance to change your past, what would you do?
3. How do you receive and make use of negative feedback as a backend developer?
4. What make you feel sense of achievement?
---
## Team Work
2. What kind of peer you don't like to work with?
3. What kind of admin you don't like to work with?
4.  How would your current co-workers describe you? 
5.  How do you share negative feedback with your co-workers?
6.  What is the difference between a good manager and a bad manager you think?
7.  What's your imagination of a good team work?
---
## Career Plan
2. Why do you want to work for us? 
3. What do you hope to have achieved in three years?
4. What would you want to do when you retire?
5.  Why you leave your last job?
---
## Jobs Experience
1. How do you approach setting professional goals?
    <details><summary>Answer</summary>
    I image myself in the future, and I think about what I want to do in the future. Then I set goals to achieve my dream.
    </details>
2. How do you prevent technical debt?
    <details><summary>Answer</summary>
    </details>
3. Describe a situation where you successfully convinced others of your ideas
4. What has been your role in development projects in the past?
    <details><summary>Answer</summary>
      Support and discuss system design and database schema design with tech lead. Build the code base for co-worker and write the test code. Support frontend to connect our API services.
    </details>
5.  How would you describe the software lifecycle at your most recent position? What did you enjoy the most? What would have you liked to change?
6.  Describe your greatest coding strength.

7.  Describe your best project. What made it successful? Were there any challenges?
8.  What is the largest web application you have ever worked on? - and what coding were you responsible for?
9.  What problem did you solve that had a significant impact on the company?
10. What the biggest challenge you have faced in your recently project? How did you solve it?
    <details><summary>Answer</summary>
    We are building a remittance service with two 3-party API services. So we need to make sure the transaction can be record correctly and be monitored.
    The challenge is how to make sure the transaction is successful, and only charge user once. And notify the receiver the result.
    The first stage it charge from user. We use an idempotency key from client side then store into DB check it every time to prevent duplicate charge. Once we got the callback from the 3-party API service, we update the transaction status and init remittance by another API service. But this callback could be triggered multiple times, so we need to make sure the transaction is not updated twice. We use idempotency key again. Also, the API is not reliable, it is unpredictable. We must set if failed if overtime. So we give this key a short time to live. Once subscribers get the message. We will check the result store in our database. If it is failed, should notice user and admin.
    </details>
11. Kafka usage
    <details><summary>Answer</summary>
    Since our system is asyc, we use kafka to handle transaction to our 3-party API services. We use kafka to store the transaction and send it to the 3-party API services. Once we got the callback, we update the transaction status and send the message to the subscribers.
    </details>


