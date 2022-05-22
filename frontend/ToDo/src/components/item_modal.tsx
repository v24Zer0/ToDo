import React from "react";
import { Pressable, Text, View } from "react-native";
import Item from "../models/item";
import List from "../models/list";

interface Props {
    item: Item;
    list: List;
}

const ItemModal: React.FC<Props> = ({ item, list }) => {
    return (
        <View>
            <Text>
                {item.task}
            </Text>
        </View>
    );
}

export default ItemModal;