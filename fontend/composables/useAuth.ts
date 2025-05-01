import { useCookie } from 'nuxt/app';

export const useAuth = () => {
    const token = useCookie('token');
    const username = useCookie('username');
    const id = useCookie('id');

    const router = useRouter();

    const isAuthenticated = computed(() => !!token.value);

    const login = (newToken: string, newUsername: string) => {
        token.value = newToken;
        username.value = newUsername;
    };

    const logout = () => {
        token.value = null;
        username.value = null;
        id.value = null;
        router.push('/')
    };

    return {
        isAuthenticated,
        login,
        logout,
    };
};