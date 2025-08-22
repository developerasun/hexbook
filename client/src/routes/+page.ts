import type { PageLoad } from './$types';
export const load: PageLoad = async ({ fetch }) => {
	let data: any = null;
	try {
		const response = await fetch('http://localhost:3010/health');
		data = await response.json();
	} catch (error) {
		return {
			message: 'server not up'
		};
	}
	return { message: data };
};
