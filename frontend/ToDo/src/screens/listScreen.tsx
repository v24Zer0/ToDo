import React from "react";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import UserLists from "../components/user_lists";
import RootStackParamList from "./rootStackParamList";
import { Button, View } from "react-native";

type Props = NativeStackScreenProps<RootStackParamList, 'List'>;

const ListScreen = ({ navigation, route }: Props) => {
    return (
        <View>
            <UserLists />
            <Button title="To Item Screen" 
                onPress={() => navigation.navigate("Item", {list: {id: "list_id", name: "list1", user_id: "user1"}})} 
            />
        </View>
    );
}

export default ListScreen;