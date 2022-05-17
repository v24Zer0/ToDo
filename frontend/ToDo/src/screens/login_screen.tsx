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
            <Button title="To List Screen" onPress={() => navigation.navigate("List")} />
        </View> 
    );
}

export default LoginScreen;