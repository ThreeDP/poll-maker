import fastify from 'fastify'
import cookie from '@fastify/cookie'
import { createPoll } from './route/create-poll'
import { getPoll } from './route/get-poll'
import { voteOnPoll } from './route/vote-on-poll'

const app = fastify()

app.register(cookie, {
    secret: "poll-maker",
    hook: 'onRequest',
})

app.register(createPoll)
app.register(getPoll)
app.register(voteOnPoll)

app.listen({ port: 5050 }).then(() =>{
    console.log("HTTP server running")
})