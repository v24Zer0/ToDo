import React from "react";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import Login from "../components/login";
import RootStackParamList from "./rootStackParamList";
import { Button, View } from "react-native";

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const LoginScreen = ({ navigation }: Props) => {
    return (
        <View>
            <Login navigate={() => navigation.reset({ index: 0, routes:[{ name: "List" }] })} />
            <Button title="Create an account" onPress={() => navigation.navigate("Signup")} />
        </View> 
    );
}

export default LoginScreen;