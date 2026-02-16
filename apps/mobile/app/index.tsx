import { AuthForm } from "@dilocash/ui/components/auth/auth-form"
import { useEffect, useState } from "react";
import { AppLoader } from "@dilocash/ui/components/app-loader";
export default function Index() {
  const [isLoaded, setIsLoaded] = useState(false);

  useEffect(() => {
    async function prepare() {
      try {
        //await getOfflineSession(); // Verifica SecureStore
        // Simula una carga mínima para evitar parpadeos
        await new Promise(resolve => setTimeout(resolve, 10000));
      } finally {
        setIsLoaded(true);
      }
    }
    prepare();
  }, []);
  if (!isLoaded) return <AppLoader subMessage="Accediendo..." isWeb={false} />;
  return (
    <>
    <AuthForm/>
    </>
  );
}
