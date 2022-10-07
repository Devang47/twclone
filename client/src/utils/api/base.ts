export const apiAddr =
	import.meta.env.MODE === 'production'
		? 'twclone.us-east-1.elasticbeanstalk.com'
		: 'localhost:5000/api';
