import React from "react";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import ItemList from "../components/item_list";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, 'Item'>;

const ItemScreen = ({ navigation, route }: Props) => {
    return (
        <ItemList list={route.params.list} />
    );
}

export default ItemScreen;