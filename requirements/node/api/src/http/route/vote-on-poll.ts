import { z } from 'zod'
import { randomUUID } from 'node:crypto'
import { prisma } from "../../lib/prisma"
import { FastifyInstance } from 'fastify'
import { request } from 'http'

export async function voteOnPoll(app: FastifyInstance) {
    app.post("/poll/:pollId/vote", async (request, reply) => {
        const voteOnPollBody = z.object({
            pollOptionId: z.string().uuid(),
        })

        const voteOnPollParams = z.object({
            pollId: z.string().uuid()
        })

        const { pollId } = voteOnPollParams.parse(request.params)
        const { pollOptionId } = voteOnPollBody.parse(request.body)

        let { sessionId } = request.cookies
        if (sessionId) {
            
        }
        if (!sessionId) {
            sessionId = randomUUID()

            reply.setCookie('sessionId', sessionId, {
                path: '/',
                maxAge: 60 * 60 * 24 * 30,
                signed: true,
                httpOnly: true,
            })
        }

        await prisma.vote.create({
            data: {
                sessionId,
                pollId,
                pollOptionId,
            }
        })

        return reply.status(201).send()
    })
}