import client from '$lib/api/client'
import { fail, redirect } from '@sveltejs/kit'

export const actions = {
    default: async ({request, params}) => {
        const spotId = params.facilityId

        const formData = await request.formData()
        const reviewContent = formData.get("reviewContent")
        const star = formData.get("star")
        console.log(reviewContent?.toString())
        if(!reviewContent || !star) {
            return fail(400, {
                ContentError: (!reviewContent)? "本文は必須です": "",
                StarError: (star && parseFloat(star.toString())  )? "星の数が不正です": ""
            })
        }

        const comment = reviewContent.toString()
        const rating = parseFloat(star.toString())

        await client.POST("/spots/{spotId}/reviews", {body:{comment, rating,spotId, userId: "a"},params: {path:{spotId}}})

        return redirect(301, `/facilities`)
    }
}
