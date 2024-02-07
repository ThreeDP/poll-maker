import fastify from 'fastify'
import { createPoll } from './route/create-poll'
import { getPoll } from './route/get-poll'

const app = fastify()

app.register(createPoll)
app.register(getPoll)

app.listen({ port: 5050 }).then(() =>{
    console.log("HTTP server running")
})