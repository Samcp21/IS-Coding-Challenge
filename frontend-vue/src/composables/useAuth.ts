import { ref } from 'vue';

const token = ref<string | null>(localStorage.getItem('jwt'));

export function useAuth() {
    const loading = ref(false);
    const errorMessage = ref('');

    const login = async (credentials: any) => {
        loading.value = true;
        errorMessage.value = '';

        try {
            const apiUrl = import.meta.env.VITE_API_URL;
            const response = await fetch(`${apiUrl}/login`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(credentials)
            });

            if (!response.ok) throw new Error('Credenciales inválidas');

            const data = await response.json();
            token.value = data.token;
            localStorage.setItem('jwt', data.token);
        } catch (error: any) {
            errorMessage.value = error.message;
        } finally {
            loading.value = false;
        }
    };

    const logout = () => {
        token.value = null;
        localStorage.removeItem('jwt');
    };

    return {
        token,
        loading,
        errorMessage,
        login,
        logout
    };
}