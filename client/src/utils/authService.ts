import createAuth0Client, { Auth0Client, type PopupLoginOptions } from '@auth0/auth0-spa-js';
import { user, isAuthenticated, popupOpen } from '$store/auth';
import config from '../auth-config';
import { goto } from '$app/navigation';

async function createClient() {
	let auth0Client = await createAuth0Client({
		domain: config.domain,
		client_id: config.clientId
	});

	return auth0Client;
}

async function loginWithPopup(client: Auth0Client, options?: PopupLoginOptions) {
	popupOpen.set(true);
	try {
		await client.loginWithPopup(options);
		user.set(await client.getUser());
		isAuthenticated.set(true);
		goto('/');
	} catch (e) {
		console.error(e);
	} finally {
		popupOpen.set(false);
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
