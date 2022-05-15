import { RouteProp, useNavigation, useRoute } from "@react-navigation/native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { StackNavigationProp } from "@react-navigation/stack";
import React from "react";
import Login from "../components/login";
import RootStackParamList from "./rootStackParamList";

// type loginScreenProps = StackNavigationProp<RootStackParamList, 'Item'>;

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const LoginScreen = ({ navigation, route }: Props) => {
    return (
        <Login />
    );
}

export default LoginScreen;