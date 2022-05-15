import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";
import ItemList from "../components/item_list";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const ItemScreen = ({ navigation, route }: Props) => {
    return (
        <ItemList list={{id: "unique_list1", name: "list1", user_id: "user1"}} />
    );
}

export default ItemScreen;