import fastify from 'fastify'

const app = fastify()

app.get("/hello", () => {
    return ("Hello NLW")
})

app.listen({ port: 5050 }).then(() =>{
    console.log("HTTP server runn")
})