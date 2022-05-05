import React from "react";
import { Text, View } from "react-native";
import Item from "../models/item";

interface ItemProps {
    item: Item;
}

const ItemComponent: React.FC<ItemProps> = ({ item }) => {
    return (
        <View>
            <Text>
                {item.task}
            </Text>
        </View>
    );
}

export default ItemComponent;