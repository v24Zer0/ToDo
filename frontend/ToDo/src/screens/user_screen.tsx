import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";
import { Text, View } from "react-native";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, 'User'>;

const UserScreen = ({ navigation, route }: Props) => {
    return (
        <View>
            <Text>
                {route.params.user.username}
            </Text>
        </View>
    );
}

export default UserScreen;