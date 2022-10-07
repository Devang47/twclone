import { socket } from '$store';
import { user } from '$store/auth';
import axios from 'axios';
import { get } from 'svelte/store';
import { apiAddr } from './base';

export const getTweets = async (token: string, limit: number = 15) => {
	const res = await axios.get('http://' + apiAddr + `/get-tweets?limit=${limit}`, {
		headers: {
			Authorization: token
		}
	});
	return res;
};

export const getTweetsByUser = async (UID: string = '12341324234') => {
	let authKey = get(user)?.uid as string;

	const res = await axios.get('http://' + apiAddr + `/get-tweets/${UID}`, {
		headers: {
			Authorization: authKey
		}
	});
	return res;
};

export const postTweets = async (token: string, tweet: Tweet) => {
	const res = await axios.post(
		'http://' + apiAddr + '/post-tweet',
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

export const deleteTweet = async (token: string, tweet: Tweet) => {
	const res = await axios.post(
		'http://' + apiAddr + '/delete-tweet',
		{
			...tweet
		},
		{
			headers: {
				Authorization: token
			}
		}
	);

	get(socket).send(JSON.stringify({ msg: 'tweets-updated', user_id: '' }));
	return res;
};

export const likeTweet = async (token: string, id: string) => {
	const res = await axios.get('http://' + apiAddr + '/like-tweet', {
		params: {
			id
		},
		headers: {
			Authorization: token
		}
	});

	return res;
};
