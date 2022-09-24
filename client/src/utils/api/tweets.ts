import { user } from '$store/auth';
import axios from 'axios';
import { get } from 'svelte/store';
import { apiAddr } from './base';

export const getTweets = async (token: string) => {
	const res = await axios.get(apiAddr + '/get-tweets?limit=15', {
		headers: {
			Authorization: token
		}
	});
	return res;
};

export const postTweets = async (token: string, tweet: Tweet) => {
	const res = await axios.post(
		apiAddr + '/post-tweet',
		{
			...tweet
		},
		{
			headers: {
				Authorization: token
			}
		}
	);
	return res;
};

export const likeTweet = async (token: string, id: string) => {
	const res = await axios.get(apiAddr + '/like-tweet', {
		params: {
			id
		},
		headers: {
			Authorization: token
		}
	});

	return res;
};
