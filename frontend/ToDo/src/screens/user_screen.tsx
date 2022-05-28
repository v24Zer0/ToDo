import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";
import { Alert, Button, Text, View } from "react-native";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, 'User'>;

const UserScreen = ({ navigation, route }: Props) => {
    return (
        <View>
            <Text>
                {route.params.user.username}
            </Text>
            <Button title="Logout" onPress={() => navigation.reset({ index: 0, routes:[{ name: "Login" }] })} />
            <Button title="Delete account" onPress={() => Alert.alert("Account deleted")} />
        </View>
    );
}

export default UserScreen;