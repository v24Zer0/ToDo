import React from "react";
import { View } from "react-native";
import Item from "../models/item";

interface Props {
    item: Item;
    setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
}

const ItemModal: React.FC<Props> = ({ item, setModalVisible }) => {
    return (
        <View>
            {item.task}
        </View>
    );
}

export default ItemModal;