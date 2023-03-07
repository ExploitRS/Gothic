import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
    const response = await fetch("api/status");
    const data = await response.json();
    if (!response.ok) {
        throw error(500, `Error: ${response.statusText}`);
    }
    const { name, text } = data;

    return {
        name,
        text
    }
}