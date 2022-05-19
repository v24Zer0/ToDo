import React from "react";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import UserLists from "../components/user_lists";
import RootStackParamList from "./rootStackParamList";
import { Button, View } from "react-native";

type Props = NativeStackScreenProps<RootStackParamList, 'List'>;

const ListScreen = ({ navigation, route }: Props) => {
    return (
        <View>
            <UserLists navigate={(list) => navigation.navigate("Item", {list})} />
        </View>
    );
}

export default ListScreen;