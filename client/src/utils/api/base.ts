export const apiAddr =
	import.meta.env.MODE === 'production'
		? 'https://twclone.awsapp.xyz/api'
		: 'http://localhost:5000/api';
