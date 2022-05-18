import React from "react";
import { Pressable, Text, View } from "react-native";
import Item from "../models/item";

interface Props {
    item: Item;
}

const ItemModal: React.FC<Props> = ({ item }) => {
    return (
        <View>
            <Text>
                {item.task}
            </Text>
        </View>
    );
}

export default ItemModal;