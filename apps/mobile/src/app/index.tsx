/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { useEffect, useState } from "react";
import { AppLoader } from "@dilocash/ui/components/app-loader";
import { KeyboardAvoidingView, Platform } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import MainScreen from "./main";
export default function Index() {
  const [isLoaded, setIsLoaded] = useState(false);
  useEffect(() => {
    async function prepare() {
      try {
        //await getOfflineSession(); // Check SecureStore
        // emulates a minimal load to avoid flickering
        await new Promise((resolve) => setTimeout(resolve, 1000));
      } finally {
        setIsLoaded(true);
      }
    }
    prepare();
  }, []);
  if (!isLoaded) return <AppLoader subMessage="Accediendo..." isWeb={false} />;
  return (
    <SafeAreaView style={{ flex: 1 }}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      >
        <MainScreen />
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
