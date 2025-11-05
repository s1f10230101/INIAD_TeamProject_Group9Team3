import client from '$lib/api/client.js'

export async function load({params}) {
    const {facilityId} = params

    return {
        reviewPromise: client.GET("/spots/{spotId}/reviews", {params: {path: {spotId: facilityId}}})
    }
}
