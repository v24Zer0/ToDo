import React from "react";
import { Text, View } from "react-native";
import Item from "../models/item";

interface ItemProps {
    item: Item;
    // setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
    // setModalItem: React.Dispatch<React.SetStateAction<Item>>;
}

// Add Pressable component to trigger modal display
// onPress triggers Modal
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