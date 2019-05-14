//
// MatchScoreBreakdown2015.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation


/** See the 2015 FMS API documentation for a description of each value */

public struct MatchScoreBreakdown2015: Codable {

    public enum Coopertition: String, Codable {
        case _none = "None"
        case unknown = "Unknown"
        case stack = "Stack"
    }
    public var blue: MatchScoreBreakdown2015Alliance?
    public var red: MatchScoreBreakdown2015Alliance?
    public var coopertition: Coopertition?
    public var coopertitionPoints: Int?

    public init(blue: MatchScoreBreakdown2015Alliance?, red: MatchScoreBreakdown2015Alliance?, coopertition: Coopertition?, coopertitionPoints: Int?) {
        self.blue = blue
        self.red = red
        self.coopertition = coopertition
        self.coopertitionPoints = coopertitionPoints
    }

    public enum CodingKeys: String, CodingKey { 
        case blue
        case red
        case coopertition
        case coopertitionPoints = "coopertition_points"
    }


}
