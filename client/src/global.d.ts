interface User {
	_id: string;
	username: string;
	email: string;
	photo: string;
	uid: string;
}

interface Likes {
	total_likes: number;
	liked_by: string[];
}

interface Tweet {
	id: string;
	author: string;
	author_uid: string;
	tweet: string;
	published_on: string;
	likes: Likes;
}
