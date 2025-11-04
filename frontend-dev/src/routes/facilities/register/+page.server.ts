import client from "$lib/api/client";
import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
  default: async ({ request }) => {
    const formData = await request.formData();
    const facilityName = formData.get("facilityName");
    const location = formData.get("location");
    // 使っていないform
    // const _hours = formData.get("hours")
    // const _priceRange = formData.get("priceRange")
    const description = formData.get("description");

    if (facilityName && location && description) {
      const { response } = await client.POST("/spots", {
        body: {
          name: facilityName.toString(),
          address: location.toString(),
          description: description.toString(),
        },
      });

      redirect(response.status, "/");
    } else {
      return fail(400, {
        facilityName,
        missingName: !facilityName,
        location,
        missingLocation: !location,
        description,
        missingDescription: !description,
      });
    }
  },
};
