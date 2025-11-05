import client from "$lib/api/client";

export async function load({ params }) {
  const facilitiId = params.facilityId;
  return {
    spotPromise: client.GET("/spots/{spotId}", {
      params: { path: { spotId: facilitiId } },
    }),
    reviewPromise: client.GET("/spots/{spotId}/reviews", {
      params: { path: { spotId: facilitiId } },
    }),
  };
}
