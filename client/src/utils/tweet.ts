import { loading, tweetsData } from '$store';
import { user } from '$store/auth';
import axios from 'axios';
import { get } from 'svelte/store';
import { getTweets, likeTweet, postTweets } from './api/tweets';

export const handleLikeTweet = async (id: string, key: string) => {
	await likeTweet(key, id);
};

export const handlePostTweet = async (e: string) => {
	try {
		loading.set(true);
		let authKey = get(user)?.uid as string;

		const dateData = (await axios.get('http://worldtimeapi.org/api/timezone/Asia/Kolkata')).data;

		dateData.datetime;

		await postTweets(authKey, {
			id: '',
			author: get(user)?.username as string,
			author_uid: get(user)?.uid as string,
			tweet: e,
			published_on: new Date(dateData.datetime).toUTCString(),
			likes: { total_likes: 0, liked_by: [] }
		});

		loading.set(false);

		return true;
	} catch (error) {
		console.log(error);
		return false;
	}
};
