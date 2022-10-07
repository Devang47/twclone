import { loading, tweetsData } from '$store';
import { user } from '$store/auth';
import { get } from 'svelte/store';
import { getTweets, likeTweet, postTweets } from './api/tweets';

export const handleLikeTweet = async (id: string, key: string) => {
	await likeTweet(key, id);
};

export const handlePostTweet = async (e: string) => {
	try {
		loading.set(true);
		let authKey = get(user)?.uid as string;

		await postTweets(authKey, {
			id: '',
			author: get(user)?.username as string,
			author_uid: get(user)?.uid as string,
			tweet: e,
			published_on: '',
			likes: { total_likes: 0, liked_by: [] }
		});

		loading.set(false);

		return true;
	} catch (error) {
		console.log(error);
		return false;
	}
};
