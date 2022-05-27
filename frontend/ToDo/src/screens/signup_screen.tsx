import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";
import { Button, View } from "react-native";
import Signup from "../components/signup";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, "Signup">;

const SignupScreen = ({ navigation }: Props) => {
    return (
        <View>
            <Signup />
            <Button title="Already have an account?" onPress={() => navigation.navigate("Login")} />
        </View>
    );
}

export default SignupScreen