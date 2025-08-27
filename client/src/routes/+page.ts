import type { PageLoad } from './$types';
import { PUBLIC_HTTP_ENDPOINT } from '$env/static/public';
export const load: PageLoad = async ({ fetch }) => {
	let data: any = null;
	try {
		const response = await fetch(PUBLIC_HTTP_ENDPOINT);
		data = await response.json();
	} catch (error) {
		return {
			message: 'server not up'
		};
	}
	return { message: data };
};
