import React from "react";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import Login from "../components/login";
import RootStackParamList from "./rootStackParamList";
import { Button, View } from "react-native";

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const LoginScreen = ({ navigation, route }: Props) => {
    return (
        <View>
            <Login />
            <Button title="Signup" onPress={() => navigation.navigate("Signup")} />
            <Button title="Login" onPress={() => navigation.reset({ index: 0, routes:[{ name:"List" }] })} />
        </View> 
    );
}

export default LoginScreen;