import { Text, View } from "react-native";
import "../global.css"
import { AuthForm } from "@dilocash/ui/components/auth/auth-form"
export default function Index() {
  return (
    <>
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text className="text-2xl text-red-500 font-bold">Edit app/index.tsx to edit this screen.</Text>
    </View>
    <View className="flex-1 items-center justify-center bg-white">
        <Text className="text-xl font-bold text-blue-500">
          Welcome to Nativewind!
        </Text>
        <AuthForm type="login"/>
      </View>
    </>
    
  );
}
