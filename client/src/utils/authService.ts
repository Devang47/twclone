import createAuth0Client, { Auth0Client, type PopupLoginOptions } from '@auth0/auth0-spa-js';
import { user, isAuthenticated } from '$store/auth';
import config from '../auth-config';
import { createUser } from '$utils/api/user';
import { loading } from '$store';

async function createClient() {
	let auth0Client = await createAuth0Client({
		domain: config.domain,
		client_id: config.clientId,
		cacheLocation: 'localstorage'
	});

	return auth0Client;
}

async function loginWithPopup(client: Auth0Client, options?: PopupLoginOptions) {
	loading.set(true);
	try {
		await client.loginWithPopup(options);
		let userData = await client.getUser();
		user.set(await createUser(userData as User));

		isAuthenticated.set(true);
		return true;
	} catch (e) {
		console.error(e);
		return false;
	} finally {
		loading.set(false);
	}
}

function logout(client: Auth0Client) {
	return client.logout();
}

const auth = {
	createClient,
	loginWithPopup,
	logout
};

export default auth;
