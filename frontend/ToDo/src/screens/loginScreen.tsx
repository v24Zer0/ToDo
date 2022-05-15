import { RouteProp, useNavigation, useRoute } from "@react-navigation/native";
import { StackNavigationProp } from "@react-navigation/stack";
import React from "react";
import Login from "../components/login";
import RootStackParamList from "./rootStackParamList";

// type loginScreenProps = StackNavigationProp<RootStackParamList, 'Item'>;

const LoginScreen = () => {
    // const navigation = useNavigation<loginScreenProps>();
    // const route = useRoute<RouteProp<RootStackParamList, 'Login'>>();

    return (
        <Login />
    );
}

export default LoginScreen;