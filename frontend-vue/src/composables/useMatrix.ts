import { ref } from 'vue';
import { useAuth } from './useAuth';

export function useMatrix() {
    const results = ref<any>(null);
    const loading = ref(false);
    const errorMessage = ref('');
    const { token, logout } = useAuth();

    const processMatrix = async (matrixString: string) => {
        loading.value = true;
        errorMessage.value = '';
        results.value = null;

        try {
            const parsedMatrix = JSON.parse(matrixString);

            if (!Array.isArray(parsedMatrix) || !Array.isArray(parsedMatrix[0])) {
                throw new Error("El formato debe ser un arreglo bidimensional válido ej: [[1,2],[3,4]]");
            }

            const apiUrl = import.meta.env.VITE_API_URL;
            const response = await fetch(`${apiUrl}/matrix`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token.value}`
                },
                body: JSON.stringify({ data: parsedMatrix })
            });

            if (response.status === 401) {
                logout();
                throw new Error('Sesión expirada. Vuelve a iniciar sesión.');
            }

            const data = await response.json();
            if (!response.ok) throw new Error(data.error || 'Error procesando la matriz');

            results.value = data;
        } catch (error: any) {
            errorMessage.value = error.message;
        } finally {
            loading.value = false;
        }
    };

    return {
        results,
        loading,
        errorMessage,
        processMatrix
    };
}