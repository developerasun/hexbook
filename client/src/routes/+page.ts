import type { PageLoad } from './$types';
export const load: PageLoad = async ({ fetch }) => {
	const response = await fetch('http://localhost:3010/health');
	const data = await response.json();
	return { message: data };
};
