// import jwt from 'jsonwebtoken';
import axios from 'axios';
import type { User as AuthUser } from '@auth0/auth0-spa-js';
import { v4 as uuidv4 } from 'uuid';
import { get } from 'svelte/store';
import { apiAddr } from './base';

export const createUser = async (data: AuthUser) => {
	let userData: User = {
		_id: '',
		username: data?.nickname || '',
		email: data?.email || '',
		photo: data?.picture || '',
		uid: uuidv4()
	};

	try {
		await fetch('http://' + apiAddr + '/create-user', {
			method: 'POST',
			mode: 'no-cors',
			credentials: 'same-origin', // include, *same-origin, omit
			headers: {
				'Content-Type': 'application/json',
				'Access-Control-Allow-Origin': '*'
			},
			redirect: 'follow', // manual, *follow, error
			referrerPolicy: 'no-referrer',
			body: JSON.stringify(userData)
		});

		const res = await getUser(userData.email);
		return res[0];
	} catch (err) {
		console.error(err);
	}
};

export const getUser = async (email: string) => {
	try {
		const res = await axios.get('http://' + apiAddr + `/get-uuid/${email}`);

		return res.data;
	} catch (err) {
		console.error(err);
	}
};

export const getUserByUid = async (uid: string) =>
	new Promise(async (resolve, reject): Promise<any> => {
		try {
			const userData = await axios.get('http://' + apiAddr + `/user/${uid}`);
			resolve(userData.data[0] as User);
		} catch (error: any) {
			if (error.response.status === 404) {
				reject({ msg: 'not-found' });
			}
			reject({ msg: 'unknown-error' });
		}
	});
