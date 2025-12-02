import client from "$lib/api/client.js";

export async function load({ params }) {
  const { data, error } = await client.GET("/spots/{spotId}", {
    params: { path: { spotId: params.facilityId } },
  });
  if (error)
    return {
      errorMsg: error.message,
    };

  return {
    spotData: data,
  };
}
