/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import FontAwesome from "@expo/vector-icons/FontAwesome";
import { useFonts } from "expo-font";
import * as SplashScreen from "expo-splash-screen";
import React, { useEffect } from "react";
import { GluestackUIProvider } from "@dilocash/ui/components/ui/gluestack-ui-provider";
import '@/global.css';
import DatabaseProvider from "./lib/database-provider";
import * as Localization from "expo-localization";
import { initI18n } from "@dilocash/i18n";
import { Stack } from "expo-router";
import { SafeAreaProvider } from "react-native-safe-area-context";
import { AuthProvider } from "@dilocash/ui/auth/provider";
import supabase from "./lib/supabase/client";

// we get the mobile language (ej. 'en', 'es')
const deviceLanguage = Localization.getLocales()[0].languageCode ?? "en";
initI18n(false, deviceLanguage);

export {
  // Catch any errors thrown by the Layout component.
  ErrorBoundary,
} from "expo-router";

export const unstable_settings = {
  // Ensure that reloading on `/modal` keeps a back button present.
  initialRouteName: "index",
};

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  const [loaded, error] = useFonts({
    SpaceMono: require("../../assets/fonts/SpaceMono-Regular.ttf"),
    ...FontAwesome.font,
  });

  // Expo Router uses Error Boundaries to catch errors in the navigation tree.
  useEffect(() => {

    if (error) {
      console.error('expo error', error);
      throw error;
    }
  }, [error]);

  useEffect(() => {
    if (loaded) {
      SplashScreen.hideAsync();
    }
  }, [loaded]);

  if (!loaded) {
    return null;
  }

  return <RootLayoutNav />;
}

function RootLayoutNav() {
  return (
    <DatabaseProvider>
      <GluestackUIProvider mode="light">
        <AuthProvider supabase={supabase}>
          <SafeAreaProvider>
            <Stack
              screenOptions={{
                headerShown: false,
                statusBarHidden: false,
                statusBarStyle: "dark",
              }}
            />
          </SafeAreaProvider>
        </AuthProvider>
      </GluestackUIProvider>
    </DatabaseProvider>
  );
}
