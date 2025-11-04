import type { PageLoad } from "./$types";
import { error } from "@sveltejs/kit";
import client from "$lib/api/client";

export const load: PageLoad = async () => {
  const { data: spotsData, error: spotsError } = await client.GET("/spots");
  if (spotsData === undefined || spotsError) {
    error(501, spotsError);
  }
  const facilitiesWithRatings = Promise.all(
    spotsData.map(async (spot) => {
      const { data: reviewsData } = await client.GET(
        "/spots/{spotId}/reviews",
        {
          params: { path: { spotId: spot.id } },
        },
      );
      if (reviewsData && reviewsData.length > 0) {
        const totalRating = reviewsData.reduce(
          (sum, review) => sum + review.rating,
          0,
        );
        const averageRating = parseFloat(
          (totalRating / reviewsData.length).toFixed(1),
        );
        const commentCount = reviewsData.length;
        return {
          spot,
          averageRating,
          commentCount,
        };
      } else {
        return {
          spot: spot,
          averageRating: 0,
          commentCount: 0,
        };
      }
    }),
  );
  return {
    facilitiesWithRatings: facilitiesWithRatings,
  };
};
