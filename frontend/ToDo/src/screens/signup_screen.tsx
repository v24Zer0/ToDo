import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";
import { Button, View } from "react-native";
import Signup from "../components/signup";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, "Signup">;

const SignupScreen = ({ navigation, route }: Props) => {
    return (
        <View>
            <Signup />
            <Button title="Login" onPress={() => navigation.navigate("Login")} />
        </View>
    );
}

export default SignupScreen