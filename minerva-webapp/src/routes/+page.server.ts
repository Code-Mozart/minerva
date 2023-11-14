import type { PageServerLoad } from './$types';
import { env } from '$env/dynamic/private';
import { error } from '@sveltejs/kit';

const BASE_API_URL = env.BASE_API_URL;

export const load = (async ({fetch}) => {
    const response = await fetch(`${BASE_API_URL}/status`);
    if (!response.ok) {
        throw error(response.status, await response.text());
    }
    const data = response.bodyUsed ? await response.json() : null;
    return { data };
}) satisfies PageServerLoad;